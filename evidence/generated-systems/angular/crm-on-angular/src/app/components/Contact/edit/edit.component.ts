import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ContactService } from '../../../services/Contact.service';
import { SubBaseComponent } from '../../Contact/sub.base.component';


@Component({
    selector: 'app-edit-contact',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditContactComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Contact';

    contactForm: FormGroup;
    contact: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ContactService,
        private fb: FormBuilder
) {
        super(http);
        this.contactForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      title: ['', Validators.required],
      email: ['', Validators.required],
      phone: ['', Validators.required],
      mobile: ['', Validators.required],
      mailingAddress: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Owner: ['', ],
      Activities: ['', ],
      Opportunities: ['', ],
      Cases: ['', ],
      Campaigns: ['', ],
      Notes: ['', ],
      EmailMessages: ['', ],
      PreferredContactMethod: ['', ]
        });
    }

    
    updateContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod): void {
        this.route.params.subscribe((params) => {

                        this.service.updateContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexContact']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getContact(params['id']).subscribe(res => {
                this.contact = res;
            });
        });
    }
}