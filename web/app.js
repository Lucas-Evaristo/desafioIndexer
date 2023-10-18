const app = Vue.createApp({
   data() {
       return {
           emailsTemplate:'',
           searchKeyword:''
       };
   },
       methods: {
           searchEmails() {

            const searchEndpoint = 'http://localhost:4090/search?term=' + this.searchKeyword;
               fetch(searchEndpoint)
               .then(response => response.json())
               .then(data => {
               this.emailsTemplate = data;
               })
               .catch(error => {
               console.error(error);
               });
           }
         }
   });

   app.component('email-item', {
      props: ['email'],
      template: `
          <div class="bg-white shadow p-4 mb-4 rounded">
              <p class="font-bold">Message ID: {{ email['Message-ID'] }}</p>
              <p><strong>Subject:</strong> {{ email.Subject }}</p>
              <p><strong>From:</strong> {{ email.From }}</p>
              <p><strong>To:</strong> {{ email.To }}</p>
              <p><strong>Date:</strong> {{ email.Date }}</p>
              <p><strong>Bcc:</strong> {{ email.Bcc }}</p>
              <p><strong>Content:</strong> {{ email.Body }}</p>
          </div>
      `
  });

    app.mount('#app');


