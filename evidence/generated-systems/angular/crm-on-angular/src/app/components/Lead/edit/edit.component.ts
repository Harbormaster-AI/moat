import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LeadService } from '../../../services/Lead.service';
import { SubBaseComponent } from '../../Lead/sub.base.component';


@Component({
    selector: 'app-edit-lead',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLeadComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Lead';

    leadForm: FormGroup;
    lead: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LeadService,
        private fb: FormBuilder
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

    
    updateLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLead(firstName, lastName, company, email, phone, converted, Organization, Owner, Activities, Campaigns, ConvertedAccount, ConvertedContact, ConvertedOpportunity, Notes, EmailMessages, Status, Source, Rating, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLead']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLead(params['id']).subscribe(res => {
                this.lead = res;
            });
        });
    }
}