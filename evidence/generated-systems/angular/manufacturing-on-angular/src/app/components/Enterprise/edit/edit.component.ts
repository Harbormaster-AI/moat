import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EnterpriseService } from '../../../services/Enterprise.service';
import { SubBaseComponent } from '../../Enterprise/sub.base.component';


@Component({
    selector: 'app-edit-enterprise',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEnterpriseComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Enterprise';

    enterpriseForm: FormGroup;
    enterprise: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EnterpriseService,
        private fb: FormBuilder
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

    
    updateEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEnterprise(name, legalName, registrationCountry, website, taxId, BusinessUnits, Plants, Suppliers, Customers, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEnterprise']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEnterprise(params['id']).subscribe(res => {
                this.enterprise = res;
            });
        });
    }
}