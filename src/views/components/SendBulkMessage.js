export default {
    name: 'SendBulkMessage',
    data() {
        return {
            phones: [''],
            message: '',
            loading: false,
            progress: {
                total: 0,
                completed: 0,
                current: '',
                status: ''
            }
        }
    },
    methods: {
        openModal() {
            $('#modalBulkMessage').modal({
                onApprove: function () {
                    return false;
                }
            }).modal('show');
        },
        addPhone() {
            this.phones.push('')
        },
        deletePhone(index) {
            this.phones.splice(index, 1)
        },
        async handleSubmit() {
            try {
                let response = await this.submitApi()
                showSuccessInfo("Bulk messages completed")
                $('#modalBulkMessage').modal('hide');
            } catch (err) {
                showErrorInfo(err)
            }
        },
        async submitApi() {
            this.loading = true;
            try {
                const payload = {
                    message: this.message,
                    phones: this.phones.filter(p => p !== '').map(p => `${p}@${window.TYPEUSER}`)
                }
                let response = await window.http.post(`/send/bulk-message`, payload)
                this.handleReset();
                return response.data;
            } catch (error) {
                if (error.response) {
                    throw new Error(error.response.data.message);
                }
                throw new Error(error.message);
            } finally {
                this.loading = false;
            }
        },
        handleReset() {
            this.phones = [''];
            this.message = '';
            this.progress = {
                total: 0,
                completed: 0,
                current: '',
                status: ''
            };
        },
        updateProgress(progress) {
            this.progress = progress;
        }
    },
    template: `
    <div class="blue card" @click="openModal()" style="cursor: pointer">
        <div class="content">
            <a class="ui blue right ribbon label">Send</a>
            <div class="header">Bulk Message</div>
            <div class="description">
                Send message to multiple numbers at once
            </div>
        </div>
    </div>

    <div class="ui small modal" id="modalBulkMessage">
        <i class="close icon"></i>
        <div class="header">
            Send Bulk Message
        </div>
        <div class="content">
            <form class="ui form">
                <div class="field">
                    <label>Phone Numbers</label>
                    <div style="display: flex; flex-direction: column; gap: 5px">
                        <div class="ui action input" v-for="(phone, index) in phones" :key="index">
                            <input type="text" placeholder="6289..." v-model="phones[index]">
                            <button class="ui button" @click="deletePhone(index)" type="button">
                                <i class="minus circle icon"></i>
                            </button>
                        </div>
                        <button class="ui button" @click="addPhone" type="button">
                            <i class="plus icon"></i> Add Phone
                        </button>
                    </div>
                </div>

                <div class="field">
                    <label>Message</label>
                    <textarea v-model="message" placeholder="Enter your message"></textarea>
                </div>

                <div class="ui indicating progress" v-if="progress.total > 0">
                    <div class="bar" :style="{width: (progress.completed/progress.total*100) + '%'}">
                        <div class="progress">{{progress.completed}}/{{progress.total}}</div>
                    </div>
                    <div class="label">{{progress.current}}</div>
                </div>
            </form>
        </div>
        <div class="actions">
            <div class="ui approve positive right labeled icon button"
                 :class="{'loading': loading}"
                 @click="handleSubmit">
                Send
                <i class="send icon"></i>
            </div>
        </div>
    </div>
    `
}
