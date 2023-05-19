<template>
	<div class="hello">
		<h1>{{ msg }}</h1>
		<h2>{{ str }}</h2>


		<el-form ref="form" :model="form" label-position="left" label-width="180px">
			<el-form-item label="文件(file):">
				<template>
					<el-col :span="10">
						<el-upload class="upload-demo" :auto-upload="false" :on-change="handleChange"
							:on-remove="handleRemove" :before-remove="beforeRemove" :file-list="fileList">
							<el-button size="small" type="primary">点击上传</el-button>
						</el-upload>

					</el-col>
				</template>
			</el-form-item>


			<el-form-item label="文件名(filename): ">
				<template>
					<el-col :span="10">
						<el-input v-model="form.text"></el-input>
					</el-col>
				</template>
			</el-form-item>

			<el-form-item>
				<el-button type="primary" @click="onSubmit">upload</el-button>
				<el-button @click="cancelForm">取消</el-button>
			</el-form-item>

		</el-form>

	</div>


</template>

<script>
import axios from 'axios'

export default {
		name: 'HelloWorld',
		props: {
			msg: String
		},
		data: function() {
			return {
				str: "文件上传",
				form: {
					file: null,
					text: '',
				},
				fileList: [],
				formData: null,
			}
		},
		mounted() {
			this.init();
		},
		methods: {
			init() {
				axios({
					method: "get",
					url: "/ping"
				}).then(res => {
					console.log(JSON.stringify(res.data))
					// this.msg = res.data.message
					this.$emit("rollback", res.data.message)
				});

				// this.$http({
				// 	method: "get",
				// 	url: "/ping"
				// }).then(res => {
				// 	alert(res.data)
				// });

				// request({
				// 	method: "get",
				// 	url: "/ping"
				// }).then(res => {
				// 	alert(res.data)
				// })

			},
			handleChange(file, fileList) {
				this.formData = new FormData();
				this.formData.append("file", file.raw);
			},
			onSubmit() {
				// let formData = new FormData();
				// alert("aa");
				let formData = this.formData;
				for (let key in this.form) {
					formData.append(key, this.form[key]);
				}
			
				axios({
					method: "post",
					url: "/fileupload",
					headers: {
						"Content-Type": "multipart/form-data"
					},
					withCredentials: true,
					data: formData
				}).then((res) => {
					this.$alert(res.data, 'result', {
						confirmButtonText: '确定',
						callback: action => {
							this.$message({
								type: 'info',
								message: `action: ${action}`
							});
						}
					});
				});
			},
			cancelForm() {
			
			},
			handleRemove(file, fileList) {
				this.formData = null;
			},
			beforeRemove(file, fileList) {
				return this.$confirm(`确定移除 ${file.name}？`);
			},
		},
		
	}
</script>

<style scoped>
	h3 {
		margin: 40px 0 0;
	}

	ul {
		list-style-type: none;
		padding: 0;
	}

	li {
		display: inline-block;
		margin: 0 10px;
	}

	a {
		color: #42b983;
	}
</style>
