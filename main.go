package main

import (
	// 	"log"
	// 	"net/http"

	"github.com/tisan-das/system-design-go-lang/behavioral_patterns"
)

func main() {
	// log.Print("hello world")

	// dbConn, err := repository.NewPostgresDBConnection("localhost", "tisan", "tisan", "tisan_db", 5432)
	// if err != nil {
	// 	log.Fatalf("Error occurred connecting to DB: %s", err)
	// 	return
	// }

	// groceryRepo := repository.NewGroceryRepo(dbConn)
	// groceryService := service.NewGroceryService(groceryRepo)

	// router := router.NewWebRouter(groceryService)
	// http.Handle("/", router)

	// http.ListenAndServe(":8080", router)

	// // ---- Design Pattern 01: Singleton
	// var singletonObj singleton.SingletonIface
	// // singletonObj = singleton.NewSingleTonThreadUnsafe()
	// singletonObj = singleton.NewSingleTonThreadSafe()

	// // // Iterative check
	// // // for i := 0; i < 10; i++ {
	// // // 	var variable singleton.SingletonThreadUnsafe
	// // // 	variable.GetInstance()
	// // // }

	// // Thread check
	// var waitGroup sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	waitGroup.Add(1)
	// 	go func() {
	// 		defer waitGroup.Done()
	// 		singletonObj.GetInstance()
	// 	}()
	// }
	// waitGroup.Wait()

	// // ---- Design Pattern 02: Prototype
	// var protoObj1, protoObj2 prototype.Prototype
	// protoObj1 = prototype.NewWindowPrototype()
	// protoObj2 = protoObj1.Clone()
	// protoObj1.Print()
	// protoObj2.Print()

	// // ------ Design Pattern 03: Builder
	// var builderStruct builder.IBuilder
	// builderStruct = builder.NewConcreteBuilder1()
	// builderStruct.BuildPart1()
	// builderStruct.BuildPart2()
	// builderStruct.BuildPart3()
	// builderStruct.GetResult()

	// // ------ Design Pattern 04: Factory Method
	// var productIface factory.IProduct
	// var creatorIface factory.ICreator

	// fmt.Println("\n--- operations on concret creator 1")
	// creatorIface = factory.NewConcreteCreator1()
	// creatorIface.FactoryMethod()
	// creatorIface.Operation()
	// productIface = creatorIface.GetProduct()
	// fmt.Printf("Product details: %+v\n", productIface)

	// fmt.Println("\n--- operations on concret creator 2")
	// creatorIface = factory.NewConcreteCreator2()
	// creatorIface.FactoryMethod()
	// creatorIface.Operation()
	// productIface = creatorIface.GetProduct()
	// fmt.Printf("Product details: %+v\n", productIface)

	// ------ Design Pattern 05: Abstract Factory
	// var widgetFactory abstractFactory.IWidgetFactory
	// fmt.Println("\n----- operations for windows widget")
	// widgetFactory = abstractFactory.NewWinWidgetFactory()
	// fmt.Printf("widget factory: %+v %+v %+v %+v\n", widgetFactory, widgetFactory.CreateWindow(),
	// 	widgetFactory.CreateButton(), widgetFactory.CreateMenu())

	// fmt.Println("\n----- operations for mac widget")
	// widgetFactory = abstractFactory.NewMacWidgetFactory()
	// fmt.Printf("widget factory: %+v %+v %+v %+v\n", widgetFactory, widgetFactory.CreateWindow(),
	// 	widgetFactory.CreateButton(), widgetFactory.CreateMenu())

	// ------ Design Pattern 06: Adapter
	// var target adapter.Target
	// fmt.Println("Using target instance")
	// target = adapter.NewTarget()
	// target.Request()
	// fmt.Println("Using adapter instance")
	// target = adapter.NewAdapter()
	// target.Request()

	// // ------ Design Pattern 06: Bridge
	// var abstraction bridge.Abstraction
	// var implementor bridge.Implementor
	// implementor = bridge.NewConcreteImplementorA()
	// abstraction = bridge.NewRefinedAbstraction(implementor)
	// abstraction.Operation()

	// // ------ Design Pattern Behavioral 01: Chain of command
	// var orderProcessor behavioral_patterns.OrderProcessor
	// var packageProcessor behavioral_patterns.PackageProcessor
	// var billingProcessor behavioral_patterns.BillingProcessor
	// var shipmentProcessor behavioral_patterns.ShipmentProcessor

	// orderProcessor.SetNext(&billingProcessor)
	// billingProcessor.SetNext(&packageProcessor)
	// packageProcessor.SetNext(&shipmentProcessor)

	// var task behavioral_patterns.Task

	// fmt.Println("--- Processing the task:")
	// orderProcessor.Execute(&task)
	// fmt.Println("--- Processing the processed task again:")
	// orderProcessor.Execute(&task)

	// // ------ Design Pattern Behavioral 02: Command
	// var tvDevice behavioral_patterns.TvDevice
	// var remoteController behavioral_patterns.RemoteController
	// remoteController.SetDevice(&tvDevice)

	// remoteController.Execute()
	// remoteController.Execute()

	// // ------ Design Pattern Behavioral 03: Interpreter
	// overallExpression := behavioral_patterns.Minus{
	// 	Expr1: &behavioral_patterns.Plus{
	// 		Expr1: &behavioral_patterns.Number{Value: 10},
	// 		Expr2: &behavioral_patterns.Number{Value: 5},
	// 	},
	// 	Expr2: &behavioral_patterns.Plus{
	// 		Expr1: &behavioral_patterns.Number{Value: 6},
	// 		Expr2: &behavioral_patterns.Number{Value: 7},
	// 	},
	// }
	// fmt.Println("Interpret Expression: ", overallExpression.Interpret())
	// fmt.Println("Evaluate Expression: ", overallExpression.Evaluate())

	// ------ Design Pattern Behavioral 04: Iterator
	// userCollection := behavioral_patterns.UserCollection{
	// 	UserArray: []behavioral_patterns.User{
	// 		behavioral_patterns.User{FirstName: "Tisan", LastName: "Das"},
	// 		behavioral_patterns.User{FirstName: "John", LastName: "Doe"},
	// 		behavioral_patterns.User{FirstName: "First", LastName: "Last"},
	// 	},
	// }
	// iterator := userCollection.CreateIterator()
	// for iterator.HasNext() {
	// 	fmt.Printf("%+v\n", iterator.GetNext().(*behavioral_patterns.User))
	// }

	// // ------ Design Pattern Behavioral 05: Memento
	// caretaker := behavioral_patterns.Caretaker{}
	// observer1 := behavioral_patterns.Observer{PrivateStr: "string1"}
	// observer2 := behavioral_patterns.Observer{PrivateStr: "string2"}

	// fmt.Printf("Before taking backup: %+v %+v\n", observer1, observer2)
	// caretaker.AppendMemento(observer1.CreateMemento())
	// caretaker.AppendMemento(observer2.CreateMemento())
	// observer1.SetMemento(caretaker.GetMementoIndex(1))
	// observer2.SetMemento(caretaker.GetMementoIndex(0))
	// fmt.Printf("After cross-restoring: %+v %+v\n", observer1, observer2)

	// ------ Design Pattern Behavioral 06: Observer
	// obs1 := behavioral_patterns.ConcreteObserver{
	// 	Id: 100,
	// }
	// obs2 := behavioral_patterns.ConcreteObserver{
	// 	Id: 100,
	// }
	// subject := behavioral_patterns.ConcreteSubject{}
	// subject.Attach(&obs1)
	// subject.Attach(&obs2)
	// fmt.Println("Triggering notification...")
	// subject.Notify()
	// subject.Detach(&obs2)
	// fmt.Println("Triggering notification...")
	// subject.Notify()

	// // ------ Design Pattern Behavioral 07: Strategy
	// var cache behavioral_patterns.Cache
	// // var algo behavioral_patterns.FifoEvictionAlgo
	// var algo behavioral_patterns.LSTEvictionAlgo
	// cache.InitCache(&algo)
	// cache.Add("key1", "value1")
	// cache.Add("key2", "value2")
	// cache.Add("key3", "value3")
	// cache.Add("key4", "value4")

	// ------ Design Pattern Behavioral 08: Visitor
	var areaCalc behavioral_patterns.AreaCalculator
	var shape behavioral_patterns.Shape
	shape = &behavioral_patterns.Square{Side: 10}
	shape.Accept(&areaCalc)
	shape = &behavioral_patterns.Circle{Radius: 10}
	shape.Accept(&areaCalc)
}
