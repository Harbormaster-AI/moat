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
                  name: ['', Validators.required],
      customerCode: ['', Validators.required],
      address: ['', Validators.required],
      Enterprises: ['', ],
      SalesOrders: ['', ],
      CustomerType: ['', ]
        });
    }

    
    updateCustomer(name, customerCode, address, Enterprises, SalesOrders, CustomerType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCustomer(name, customerCode, address, Enterprises, SalesOrders, CustomerType, params['id'])
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