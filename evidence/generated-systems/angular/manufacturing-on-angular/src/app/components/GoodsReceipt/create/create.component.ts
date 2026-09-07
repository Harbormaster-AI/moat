import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { GoodsReceiptService } from '../../../services/GoodsReceipt.service';
import { GoodsReceipt } from '../../../models/GoodsReceipt';
import { SubBaseComponent } from '../../GoodsReceipt/sub.base.component';

@Component({
    selector: 'app-create-goodsReceipt',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateGoodsReceiptComponent extends SubBaseComponent implements OnInit {

    title = 'Add GoodsReceipt';

    goodsReceiptForm: FormGroup;
    goodsReceipt: GoodsReceipt;

    constructor( http: HttpClient,
        private goodsReceiptService: GoodsReceiptService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.goodsReceiptForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  receiptNumber: ['', Validators.required],
      receiptDate: ['', Validators.required],
      PurchaseOrder: ['', ],
      Warehouse: ['', ],
      Lines: ['', ],
      Status: ['', ]
        });
    }

    
    addGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status): void {
        this.goodsReceiptService
        .addGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status)
            .subscribe(() => {
                this.router.navigate(['/indexGoodsReceipt']);
            });
    }

    ngOnInit(): void {
    }
}