import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PlannedOrderService } from '../../../services/PlannedOrder.service';
import { PlannedOrder } from '../../../models/PlannedOrder';
import { SubBaseComponent } from '../../PlannedOrder/sub.base.component';

@Component({
    selector: 'app-create-plannedOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePlannedOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add PlannedOrder';

    plannedOrderForm: FormGroup;
    plannedOrder: PlannedOrder;

    constructor( http: HttpClient,
        private plannedOrderService: PlannedOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.plannedOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  plannedOrderNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      dueDate: ['', Validators.required],
      MrpRun: ['', ],
      Item: ['', ],
      Plant: ['', ],
      OrderType: ['', ],
      Status: ['', ]
        });
    }

    
    addPlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status): void {
        this.plannedOrderService
        .addPlannedOrder(plannedOrderNumber, quantity, dueDate, MrpRun, Item, Plant, OrderType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexPlannedOrder']);
            });
    }

    ngOnInit(): void {
    }
}