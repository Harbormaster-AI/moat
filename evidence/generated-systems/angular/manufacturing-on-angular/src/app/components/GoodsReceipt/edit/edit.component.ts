import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { GoodsReceiptService } from '../../../services/GoodsReceipt.service';
import { SubBaseComponent } from '../../GoodsReceipt/sub.base.component';


@Component({
    selector: 'app-edit-goodsReceipt',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditGoodsReceiptComponent extends SubBaseComponent implements OnInit {

    title = 'Edit GoodsReceipt';

    goodsReceiptForm: FormGroup;
    goodsReceipt: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: GoodsReceiptService,
        private fb: FormBuilder
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

    
    updateGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateGoodsReceipt(receiptNumber, receiptDate, PurchaseOrder, Warehouse, Lines, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexGoodsReceipt']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getGoodsReceipt(params['id']).subscribe(res => {
                this.goodsReceipt = res;
            });
        });
    }
}