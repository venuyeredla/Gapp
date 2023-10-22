function createWorld(){
    return (...args) =>{
        return "Hello world"
    }
}
var numer=0
function cc(n){
    numer=parseInt(n)
    return ()=> {
        return numer++
    }
}

//const f=createWorld();
const f=cc(10);
for(let i=0;i<3;i++){
    console.log(f());
}

var expect = function(val) {
    return {
        toBe:(val1)=>{
            if(val===val1){
                return true
            }else{
                throw "Not Equal"
            }
        },
        notToBe:(val2)=>{
             if(val !==val2){
                return true
            }else{
                throw "Equal"
            }
        }
    }
};

function greeting(lang) {
    console.log(`${lang}: I am ${this.name}`);
  }
  const john = {
    name: 'John'
  };
  const jane = {
    name: 'Jane'
  };
  const greetingJohn = greeting.bind(john, 'en');
  greetingJohn();
  const greetingJane = greeting.bind(jane, 'es');
  greetingJane();