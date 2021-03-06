# Golang Design Patterns
![alt](https://github.com/MrVWY/GolangDesignPatterns/blob/master/g.jpg)
----
* 创建模式(Creational Patterns)
    * 抽象工厂  
    * [生成器](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Creational%20Patterns/builder.go): 使用简单对象构建复杂对象  
    * 工厂方法  
    * 对象池  
    * [单例](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Creational%20Patterns/Singleton.go): 将类型的实例化限制为一个对象  

* 结构模式(Structural Patterns)
    * [适配器](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/adapter.go): 适配另一个不兼容的接口来一起工作  
    * [桥接](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/bridging.go): 将接口与其实现分离，以便两者可以独立变化  
    * 组合  
    * [装饰](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/decorator.go): 静态或动态地向对象添加行为  
    * [外观](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/facede.go): 使用一种类型作为许多其他类型的API  
    * [享元](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/Flyweight.go): 运用共享技术有效地支持大量细粒度的对象  
    * [代理](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Structural%20Patterns/proxy.go): 运用共享技术有效地支持大量细粒度的对象  

* 行为模式(Behavioral Patterns)
    * [观察者](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/observer.go)  
    * [职责链](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/chain_of_reponsibility.go)
    * [中介者](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/mediator.go) :对象与对象之间存在大量的关联关系，这样势必会导致系统的结构变得很复杂，同时若一个对象发生改变，我们也需要跟踪与之相关联的对象，同时作出相应的处理  
    * [备忘录](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/memorandum.go) :在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态  
    * [状态](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/state.go)
    * [策略](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Behavioral%20Patterns/strategy.go)
    * 模板  
    * 访问者

* 并发模式(Concurrency Patterns)
    * [Active Object]() 

* 函数式选项模式(Functional Options Pattern)
    * [函数式选项式](https://github.com/MrVWY/GolangDesignPatterns/blob/master/Functional%20Options%20Patterns/Options.go)