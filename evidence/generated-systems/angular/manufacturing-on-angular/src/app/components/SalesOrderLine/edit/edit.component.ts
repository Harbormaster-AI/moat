import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SalesOrderLineService } from '../../../services/SalesOrderLine.service';
import { SubBaseComponent } from '../../SalesOrderLine/sub.base.component';


@Component({
    selector: 'app-edit-salesOrderLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSalesOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SalesOrderLine';

    salesOrderLineForm: FormGroup;
    salesOrderLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SalesOrderLineService,
        private fb: FormBuilder
) {
        super(http);
        this.salesOrderLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      unitPrice: ['', Validators.required],
      dueDate: ['', Validators.required],
      SalesOrder: ['', ],
      Item: ['', ]
        });
    }

    
    updateSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSalesOrderLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSalesOrderLine(params['id']).subscribe(res => {
                this.salesOrderLine = res;
            });
        });
    }
}