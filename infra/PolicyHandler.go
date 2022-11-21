
forEach: BoundedContext
fileName: PolicyHandler.go
path: {{name}}/{{name}}
---
package {{name}}

{{#policyExists policies}}
import (
	"github.com/mitchellh/mapstructure"
)
{{/policyExists}}

{{#policies}}
{{#relationEventInfo}}
func whenever{{eventValue.namePascalCase}}_{{../namePascalCase}}(data map[string]interface{}){
	
	event := New{{eventValue.namePascalCase}}()
	mapstructure.Decode(data,&event)
	{{#../../aggregates}}
	{{nameCamelCase}} := &{{namePascalCase}}{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	{{nameCamelCase}}repository.save({{nameCamelCase}})
	{{/../../aggregates}}

	// Sample Logic //
	{{#../aggregateList}}
	{{nameCamelCase}}.{{../../namePascalCase}}(event);
	{{/../aggregateList}}
}

{{/relationEventInfo}}
{{/policies}}

<function>
	window.$HandleBars.registerHelper('policyExists', function (policies, options) {
		if(Object.values(policies) != ""){
			return options.fn(this)
        }
        else{
            return options.inverse(this)
        }
		
	});
</function>