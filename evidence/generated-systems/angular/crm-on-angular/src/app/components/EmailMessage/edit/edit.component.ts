import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EmailMessageService } from '../../../services/EmailMessage.service';
import { SubBaseComponent } from '../../EmailMessage/sub.base.component';


@Component({
    selector: 'app-edit-emailMessage',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEmailMessageComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EmailMessage';

    emailMessageForm: FormGroup;
    emailMessage: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EmailMessageService,
        private fb: FormBuilder
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

    
    updateEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEmailMessage(subject, body, sentAt, messageId, Organization, Owner, Account, Contact, Lead, Case, Opportunity, Campaign, Direction, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEmailMessage']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEmailMessage(params['id']).subscribe(res => {
                this.emailMessage = res;
            });
        });
    }
}