<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>xorm转换工具</title>
  <meta name="renderer" content="webkit">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <link rel="stylesheet" href="https://www.layuicdn.com/layui/css/layui.css"  media="all">
  <!-- 注意：如果你直接复制所有代码到本地，上述css路径需要改成你本地的 -->
</head>

<body>
  <!-- 让IE8/9支持媒体查询，从而兼容栅格 -->
  <!--[if lt IE 9]>
    <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
    <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
  <![endif]-->     

  <fieldset class="layui-elem-field layui-field-title" style="margin-top: 20px;">
    <legend>xorm 转换</legend>
  </fieldset>
  <div class="layui-bg-gray" style="padding: 30px;">
    <div class="layui-inline layui-col-md3" style="padding: 50px 20px;">
      <input name="" placeholder="xorm实例名,如GlobalEngine" class="layui-input" value="GlobalEngine" id="engineName"/>
    </div>
    <div class="layui-inline layui-col-md3" style="padding: 50px 20px;">
      <button type="button" class="layui-btn" id="convert">转换</button>
    </div>
    
    <div class="layui-row layui-col-space15">
      <div class="layui-col-md6">
        <div class="layui-panel">

          <div style="padding: 0px 30px;">
            <div style="padding: 0px 0px 10px 0px;">sql</div>
            <div style="padding: 0px 0px 10px 0px;">
              <textarea placeholder="请输入内容" rows="20" class="layui-textarea" id="sql">CREATE TABLE IF NOT EXISTS `runoob_tbl`(
   `runoob_id` INT UNSIGNED AUTO_INCREMENT,
   `runoob_title` VARCHAR(100) NOT NULL,
   `runoob_author` VARCHAR(40) NOT NULL,
   `submission_date` DATE,
   PRIMARY KEY ( `runoob_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;</textarea>
            </div>
          </div>
        </div>   
      </div>
      <div class="layui-col-md6">
        <div class="layui-panel">
          <div style="padding: 0px 30px;">
            <div style="padding: 0px 0px 10px 0px;" >结果</div>
            <div style="padding: 0px 0px 10px 0px;">
              <textarea placeholder="请输入内容" rows="20" class="layui-textarea" id="result"></textarea>
            </div>
          </div>
        </div>   
      </div>
    </div> 
  </div>

  <script src="https://www.layuicdn.com/layui/layui.js" charset="utf-8"></script>
  <!-- 注意：如果你直接复制所有代码到本地，上述 JS 路径需要改成你本地的 -->
  <script>
      var dropdown = layui.dropdown
      ,util = layui.util
      ,layer = layui.layer
      ,$ = layui.jquery;
      layui.use(['dropdown', 'util', 'layer'], function(){
          /* 按钮事件 */
          $("#convert").click(function(){
            $.ajax({
                url:"/sql2go/xormConvert",
                type:"POST",
                contentType : 'application/json',
                data: JSON.stringify({"sql": $("#sql").val(), "engineName": $("#engineName").val()}),
                success:function(datas){
                  var errMsg = datas["err_msg"]
                  if (errMsg != undefined){
                    layer.msg("错误:"+errMsg);
                    return
                  }
                  $("#result").val(datas["result"]);
                },
                error: function() {
                    layer.msg('修改失败');
                }
            });
          });
    });
  </script>
</body>