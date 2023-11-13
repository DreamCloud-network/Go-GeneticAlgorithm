package basicmachinery

import (
	"log"
	"time"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/evoso/environment"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

// Represents the basic struct of any automata active receptor organelle.
// This organelle is responsible for perceive and get specifics things form outside to inside.
type collector struct {
	Organelle

	colletorThingType environment.ThingType
}

func newCollectorGene(thingsType environment.ThingType) *genes.Gene {
	newCollectorGene := genes.NewGene()

	newCollectorGene.AddCodon(COLECTOR_CODON)
	newCollectorGene.AddCodon(SEPARATOR_CODON)

	newCollectorGene.AddCodons(codons.UintToCodons(100))

	newCollectorGene.AddCodon(SEPARATOR_CODON)

	newCollectorGene.AddCodons(codons.UintToCodons(uint(thingsType)))

	return newCollectorGene
}

func decodeCollectorGene(gene *genes.Gene) (int, environment.ThingType, error) {
	decodeState := 0

	geneCodons := gene.ReadCode()
	idleSleepTime := 0
	thingType := environment.Unknown

	for {
		switch decodeState {
		case 0:
			// Verify if the gene has a valid collector gene init
			if len(geneCodons) < 4 {
				log.Println("basicunit.collector.decodeCollectorGene - Invalid gene code")
				return 0, environment.Unknown, ErrInvalidGeneCode
			}

			if (geneCodons[0] != codons.INIT_CODON) || (geneCodons[1] != COLECTOR_CODON) || (geneCodons[2] != SEPARATOR_CODON) {
				log.Println("basicunit.collector.decodeCollectorGene - Invalid gene code")
				return 0, environment.Unknown, ErrInvalidGeneCode
			}
			geneCodons = geneCodons[3:]
			decodeState++
		case 1:
			idleSleepTimeCodons := make([]codons.Codon, 0)

			for pos, codon := range geneCodons {
				if codon == SEPARATOR_CODON {
					if len(idleSleepTimeCodons) == 0 {
						log.Println("basicunit.collector.decodeCollectorGene - Invalid gene code")
						return 0, environment.Unknown, ErrInvalidGeneCode
					} else {
						decodedSleepTime, err := codons.CodonsToUint(idleSleepTimeCodons)
						if err != nil {
							log.Println("basicunit.collector.decodeCollectorGene - Error reading sleep time codons")
							return 0, environment.Unknown, err
						}
						idleSleepTime = int(decodedSleepTime)

						geneCodons = geneCodons[pos+1:]

						decodeState++
					}
					break
				}
				idleSleepTimeCodons = append(idleSleepTimeCodons, codon)
			}
		case 2:
			if len(geneCodons) <= 0 {
				log.Println("basicunit.collector.decodeCollectorGene - Invalid gene code")
				return 0, environment.Unknown, ErrInvalidGeneCode
			}
			thingTypeDecoded, err := codons.CodonsToUint(geneCodons)
			if err != nil {
				log.Println("basicunit.collector.decodeCollectorGene - Error reading thing type codons")
				return 0, environment.Unknown, err
			}
			thingType = environment.ThingType(thingTypeDecoded)
			return idleSleepTime, thingType, nil

		}
	}
}

func NewCollector(basicUnit *BasicMachinery, gene *genes.Gene) error {

	// Decode genes to get parameters
	idleSleepTime, thingType, err := decodeCollectorGene(gene)
	if err != nil {
		log.Println("basicunit.collector.NewCollector - Error decoding gene")
		return err
	}

	// Create the new collector
	newCollector := &collector{
		Organelle:         *NewOrganelle(environment.Collector, basicUnit),
		colletorThingType: thingType,
	}

	newCollector.SetObject(newCollector)

	newCollector.IdleSleepTime = idleSleepTime

	// Add collector in the internal environment of the basic unit
	basicUnit.GetInternalEnvironment().AddThing(&newCollector.Thing)

	// Activate the collector
	return newCollector.Activate()
}

func (collector *collector) Activate() error {
	// The active receptor must have two environments, the internal and the external
	// Testing the external environment
	if collector.Machinary.GetExternalEnvironment() == nil {
		log.Println("basicunit.activeReceptor.Activate - No external environment")
		return ErrNoExternalEnv
	}

	// Testing the internal environment
	if collector.GetExternalEnvironment() == nil {
		log.Println("basicunit.activeReceptor.Activate - No internal environment")
		return ErrNoInternalEnv
	}

	go collector.Run()

	return nil
}

func (collector *collector) Run() {
	collector.Machinary.Add(1)
	defer collector.Machinary.Done()

	for collector.Machinary.alive {
		// Get one available time quantum to execute job.
		timeQuantum := collector.GetExternalEnvironment().GetOneThingType(environment.TimeQuantum)
		if timeQuantum != nil {

			// Perceive the things from the external environment
			thing := collector.Machinary.GetExternalEnvironment().GetOneThingType(collector.colletorThingType)

			if thing != nil {
				// Put in the internal environment
				collector.GetExternalEnvironment().AddThing(thing)

				// Give back the time quantum
				collector.GetExternalEnvironment().AddThing(timeQuantum)
			} else {
				// Give back the time quantum
				collector.GetExternalEnvironment().AddThing(timeQuantum)

				// Sleep for a while
				time.Sleep(time.Duration(collector.IdleSleepTime) * time.Millisecond)
			}
		}
	}
}
