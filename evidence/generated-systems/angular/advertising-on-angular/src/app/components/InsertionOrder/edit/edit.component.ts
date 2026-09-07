import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsertionOrderService } from '../../../services/InsertionOrder.service';
import { SubBaseComponent } from '../../InsertionOrder/sub.base.component';


@Component({
    selector: 'app-edit-insertionOrder',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsertionOrderComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InsertionOrder';

    insertionOrderForm: FormGroup;
    insertionOrder: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsertionOrderService,
        private fb: FormBuilder
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

    
    updateInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsertionOrder(ioNumber, agreedBudget, flight, Advertiser, Agency, Publisher, Campaigns, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsertionOrder']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsertionOrder(params['id']).subscribe(res => {
                this.insertionOrder = res;
            });
        });
    }
}