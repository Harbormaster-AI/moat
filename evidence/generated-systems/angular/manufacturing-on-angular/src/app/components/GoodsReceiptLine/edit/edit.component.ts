import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';
import { SubBaseComponent } from '../../GoodsReceiptLine/sub.base.component';


@Component({
    selector: 'app-edit-goodsReceiptLine',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditGoodsReceiptLineComponent extends SubBaseComponent implements OnInit {

    title = 'Edit GoodsReceiptLine';

    goodsReceiptLineForm: FormGroup;
    goodsReceiptLine: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: GoodsReceiptLineService,
        private fb: FormBuilder
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

    
    updateGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction): void {
        this.route.params.subscribe((params) => {

                        this.service.updateGoodsReceiptLine(lineNumber, receivedQuantity, acceptedQuantity, rejectedQuantity, lot, GoodsReceipt, Item, InventoryTransaction, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexGoodsReceiptLine']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getGoodsReceiptLine(params['id']).subscribe(res => {
                this.goodsReceiptLine = res;
            });
        });
    }
}