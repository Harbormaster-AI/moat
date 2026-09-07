import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OrganizationService } from '../../../services/Organization.service';
import { Organization } from '../../../models/Organization';
import { SubBaseComponent } from '../../Organization/sub.base.component';

@Component({
    selector: 'app-create-organization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOrganizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Organization';

    organizationForm: FormGroup;
    organization: Organization;

    constructor( http: HttpClient,
        private organizationService: OrganizationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.organizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      defaultCurrency: ['', Validators.required],
      defaultLocale: ['', Validators.required],
      website: ['', Validators.required],
      Users: ['', ],
      Accounts: ['', ],
      Teams: ['', ],
      Territories: ['', ],
      Products: ['', ],
      PriceBooks: ['', ],
      Campaigns: ['', ]
        });
    }

    
    addOrganization(name, defaultCurrency, defaultLocale, website, Users, Accounts, Teams, Territories, Products, PriceBooks, Campaigns): void {
        this.organizationService
        .addOrganization(name, defaultCurrency, defaultLocale, website, Users, Accounts, Teams, Territories, Products, PriceBooks, Campaigns)
            .subscribe(() => {
                this.router.navigate(['/indexOrganization']);
            });
    }

    ngOnInit(): void {
    }
}