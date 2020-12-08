package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/router"
	"html/template"
)

const (
	panicText = `<script src="http://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/viz.js"></script>
<script src="http://mnur-prod-public.oss-cn-beijing.aliyuncs.com/0/tech/full.render.js"></script>
<script type="dot" id="dotscript">
{{.}}
</script>
<script>
  window.onload=function(e){
    var viz = new Viz();
    viz.renderSVGElement(document.getElementById('dotscript').innerText)
    .then(function(element) {
      document.body.appendChild(element);
    })
    .catch(error => {
      // Create a new Viz instance (@see Caveats page for more info)
      viz = new Viz();
      // Possibly display the error
      console.error(error);
    });
  }
</script>
`
)

var panicHTMLTemplate = template.Must(template.New("PanicPage").Parse(panicText))

func UseViz(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded graphViz endpoint.")

	router.GET("/actuator/graph", func(ctx *context.HttpContext) {
		graphType := ctx.Input.QueryDefault("type", "data")
		graphString := ctx.RequiredServices.GetGraph()

		if graphType == "data" {
			ctx.Text(200, graphString)
		} else {
			ctx.Output.Header(context.HeaderContentType, context.MIMETextHTMLCharsetUTF8)
			ctx.Output.SetStatus(200)
			_ = panicHTMLTemplate.Execute(ctx.Output.GetWriter(), template.HTML(graphString))
		}
	})
}
