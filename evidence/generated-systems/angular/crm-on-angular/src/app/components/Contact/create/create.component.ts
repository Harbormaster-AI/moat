import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ContactService } from '../../../services/Contact.service';
import { Contact } from '../../../models/Contact';
import { SubBaseComponent } from '../../Contact/sub.base.component';

@Component({
    selector: 'app-create-contact',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateContactComponent extends SubBaseComponent implements OnInit {

    title = 'Add Contact';

    contactForm: FormGroup;
    contact: Contact;

    constructor( http: HttpClient,
        private contactService: ContactService,
        private fb: FormBuilder,
        private router: Router
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

    
    addContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod): void {
        this.contactService
        .addContact(firstName, lastName, title, email, phone, mobile, mailingAddress, Organization, Account, Owner, Activities, Opportunities, Cases, Campaigns, Notes, EmailMessages, PreferredContactMethod)
            .subscribe(() => {
                this.router.navigate(['/indexContact']);
            });
    }

    ngOnInit(): void {
    }
}