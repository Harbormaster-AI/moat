import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SupplierService } from '../../../services/Supplier.service';
import { SubBaseComponent } from '../../Supplier/sub.base.component';


@Component({
    selector: 'app-edit-supplier',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSupplierComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Supplier';

    supplierForm: FormGroup;
    supplier: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SupplierService,
        private fb: FormBuilder
) {
        super(http);
        this.supplierForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      supplierCode: ['', Validators.required],
      address: ['', Validators.required],
      Enterprises: ['', ],
      Items: ['', ],
      PurchaseOrders: ['', ],
      SupplierTier: ['', ],
      PaymentTerms: ['', ]
        });
    }

    
    updateSupplier(name, supplierCode, address, Enterprises, Items, PurchaseOrders, SupplierTier, PaymentTerms): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSupplier(name, supplierCode, address, Enterprises, Items, PurchaseOrders, SupplierTier, PaymentTerms, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSupplier']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSupplier(params['id']).subscribe(res => {
                this.supplier = res;
            });
        });
    }
}