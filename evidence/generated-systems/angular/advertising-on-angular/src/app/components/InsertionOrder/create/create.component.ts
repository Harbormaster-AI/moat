import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsertionOrderService } from '../../../services/InsertionOrder.service';
import { InsertionOrder } from '../../../models/InsertionOrder';
import { SubBaseComponent } from '../../InsertionOrder/sub.base.component';

@Component({
    selector: 'app-create-insertionOrder',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsertionOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Add InsertionOrder';

    insertionOrderForm: FormGroup;
    insertionOrder: InsertionOrder;

    constructor( http: HttpClient,
        private insertionOrderService: InsertionOrderService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.insertionOrderForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  ioNumber: ['', Validators.required],
      agreedBudget: ['', Validators.required],
      flight: ['', Validators.required],
      Advertiser: ['', ],
      Agency: ['', ],
      Publisher: ['', ],
      Campaigns: ['', ],
      Status: ['', ]
        });
    }

    
    addInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status): void {
        this.insertionOrderService
        .addInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status)
            .subscribe(() => {
                this.router.navigate(['/indexInsertionOrder']);
            });
    }

    ngOnInit(): void {
    }
}