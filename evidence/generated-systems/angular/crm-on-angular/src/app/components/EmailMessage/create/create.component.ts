import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EmailMessageService } from '../../../services/EmailMessage.service';
import { EmailMessage } from '../../../models/EmailMessage';
import { SubBaseComponent } from '../../EmailMessage/sub.base.component';

@Component({
    selector: 'app-create-emailMessage',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEmailMessageComponent extends SubBaseComponent implements OnInit {

    title = 'Add EmailMessage';

    emailMessageForm: FormGroup;
    emailMessage: EmailMessage;

    constructor( http: HttpClient,
        private emailMessageService: EmailMessageService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.emailMessageForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  subject: ['', Validators.required],
      body: ['', Validators.required],
      sentAt: ['', Validators.required],
      messageId: ['', Validators.required],
      Organization: ['', ],
      Owner: ['', ],
      Account: ['', ],
      Contact: ['', ],
      Lead: ['', ],
      Case: ['', ],
      Opportunity: ['', ],
      Campaign: ['', ],
      Direction: ['', ],
      Status: ['', ]
        });
    }

    
    addEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status): void {
        this.emailMessageService
        .addEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status)
            .subscribe(() => {
                this.router.navigate(['/indexEmailMessage']);
            });
    }

    ngOnInit(): void {
    }
}