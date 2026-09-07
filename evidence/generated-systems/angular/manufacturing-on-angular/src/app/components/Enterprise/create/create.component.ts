import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EnterpriseService } from '../../../services/Enterprise.service';
import { Enterprise } from '../../../models/Enterprise';
import { SubBaseComponent } from '../../Enterprise/sub.base.component';

@Component({
    selector: 'app-create-enterprise',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEnterpriseComponent extends SubBaseComponent implements OnInit {

    title = 'Add Enterprise';

    enterpriseForm: FormGroup;
    enterprise: Enterprise;

    constructor( http: HttpClient,
        private enterpriseService: EnterpriseService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.enterpriseForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      legalName: ['', Validators.required],
      registrationCountry: ['', Validators.required],
      website: ['', Validators.required],
      taxId: ['', Validators.required],
      BusinessUnits: ['', ],
      Plants: ['', ],
      Suppliers: ['', ],
      Customers: ['', ]
        });
    }

    
    addEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers): void {
        this.enterpriseService
        .addEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers)
            .subscribe(() => {
                this.router.navigate(['/indexEnterprise']);
            });
    }

    ngOnInit(): void {
    }
}