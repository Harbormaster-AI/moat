import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LeadService } from '../../../services/Lead.service';
import { Lead } from '../../../models/Lead';
import { SubBaseComponent } from '../../Lead/sub.base.component';

@Component({
    selector: 'app-create-lead',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLeadComponent extends SubBaseComponent implements OnInit {

    title = 'Add Lead';

    leadForm: FormGroup;
    lead: Lead;

    constructor( http: HttpClient,
        private leadService: LeadService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.leadForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      company: ['', Validators.required],
      email: ['', Validators.required],
      phone: ['', Validators.required],
      converted: ['', Validators.required],
      Organization: ['', ],
      Owner: ['', ],
      Activities: ['', ],
      Campaigns: ['', ],
      ConvertedAccount: ['', ],
      ConvertedContact: ['', ],
      ConvertedOpportunity: ['', ],
      Notes: ['', ],
      EmailMessages: ['', ],
      Status: ['', ],
      Source: ['', ],
      Rating: ['', ]
        });
    }

    
    addLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating): void {
        this.leadService
        .addLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating)
            .subscribe(() => {
                this.router.navigate(['/indexLead']);
            });
    }

    ngOnInit(): void {
    }
}