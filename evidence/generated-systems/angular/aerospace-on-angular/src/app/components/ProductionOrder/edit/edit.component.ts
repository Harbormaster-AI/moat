import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProductionOrderService } from '../../../services/ProductionOrder.service';
import { SubBaseComponent } from '../../ProductionOrder/sub.base.component';


@Component({
    selector: 'app-edit-productionOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProductionOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ProductionOrder';

    productionOrderForm: FormGroup;
    productionOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProductionOrderService,
        private fb: FormBuilder
) {
        super(http);
        this.productionOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  orderNumber: ['', Validators.required],
      Variant: ['', ],
      Plant: ['', ],
      AircraftOrder: ['', ],
      Status: ['', ]
        });
    }

    
    updateProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProductionOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProductionOrder(params['id']).subscribe(res => {
                this.productionOrder = res;
            });
        });
    }
}