import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CustomerService } from '../../../services/Customer.service';
import { SubBaseComponent } from '../../Customer/sub.base.component';


@Component({
    selector: 'app-edit-customer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCustomerComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Customer';

    customerForm: FormGroup;
    customer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CustomerService,
        private fb: FormBuilder
) {
        super(http);
        this.customerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      organizationName: ['', Validators.required],
      taxId: ['', Validators.required],
      dateOfBirth: ['', Validators.required],
      primaryAddress: ['', Validators.required],
      Applications: ['', ],
      Policies: ['', ],
      Claims: ['', ],
      Agents: ['', ],
      Beneficiaries: ['', ],
      CustomerType: ['', ]
        });
    }

    
    updateCustomer(firstName, lastName, organizationName, taxId, dateOfBirth, primaryAddress, Applications, Policies, Claims, Agents, Beneficiaries, CustomerType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCustomer(firstName, lastName, organizationName, taxId, dateOfBirth, primaryAddress, Applications, Policies, Claims, Agents, Beneficiaries, CustomerType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCustomer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCustomer(params['id']).subscribe(res => {
                this.customer = res;
            });
        });
    }
}