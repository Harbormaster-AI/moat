import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProductionOrderService } from '../../../services/ProductionOrder.service';
import { ProductionOrder } from '../../../models/ProductionOrder';
import { SubBaseComponent } from '../../ProductionOrder/sub.base.component';

@Component({
    selector: 'app-create-productionOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProductionOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add ProductionOrder';

    productionOrderForm: FormGroup;
    productionOrder: ProductionOrder;

    constructor( http: HttpClient,
        private productionOrderService: ProductionOrderService,
        private fb: FormBuilder,
        private router: Router
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

    
    addProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status): void {
        this.productionOrderService
        .addProductionOrder(orderNumber, Variant, Plant, AircraftOrder, Status)
            .subscribe(() => {
                this.router.navigate(['/indexProductionOrder']);
            });
    }

    ngOnInit(): void {
    }
}