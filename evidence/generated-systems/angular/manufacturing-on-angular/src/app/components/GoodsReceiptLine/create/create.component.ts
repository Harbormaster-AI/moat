import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';
import { GoodsReceiptLine } from '../../../models/GoodsReceiptLine';
import { SubBaseComponent } from '../../GoodsReceiptLine/sub.base.component';

@Component({
    selector: 'app-create-goodsReceiptLine',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateGoodsReceiptLineComponent extends SubBaseComponent implements OnInit {

    title = 'Add GoodsReceiptLine';

    goodsReceiptLineForm: FormGroup;
    goodsReceiptLine: GoodsReceiptLine;

    constructor( http: HttpClient,
        private goodsReceiptLineService: GoodsReceiptLineService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.goodsReceiptLineForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      receivedQuantity: ['', Validators.required],
      acceptedQuantity: ['', Validators.required],
      rejectedQuantity: ['', Validators.required],
      lot: ['', Validators.required],
      GoodsReceipt: ['', ],
      Item: ['', ],
      InventoryTransaction: ['', ]
        });
    }

    
    addGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction): void {
        this.goodsReceiptLineService
        .addGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction)
            .subscribe(() => {
                this.router.navigate(['/indexGoodsReceiptLine']);
            });
    }

    ngOnInit(): void {
    }
}