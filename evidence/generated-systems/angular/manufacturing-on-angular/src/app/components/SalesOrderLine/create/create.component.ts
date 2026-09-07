import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SalesOrderLineService } from '../../../services/SalesOrderLine.service';
import { SalesOrderLine } from '../../../models/SalesOrderLine';
import { SubBaseComponent } from '../../SalesOrderLine/sub.base.component';

@Component({
    selector: 'app-create-salesOrderLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSalesOrderLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add SalesOrderLine';

    salesOrderLineForm: FormGroup;
    salesOrderLine: SalesOrderLine;

    constructor( http: HttpClient,
        private salesOrderLineService: SalesOrderLineService,
        private fb: FormBuilder,
        private router: Router
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

    
    addSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item): void {
        this.salesOrderLineService
        .addSalesOrderLine(lineNumber, quantity, unitPrice, dueDate, SalesOrder, Item)
            .subscribe(() => {
                this.router.navigate(['/indexSalesOrderLine']);
            });
    }

    ngOnInit(): void {
    }
}