
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GoodsReceiptService } from '../../../services/GoodsReceipt.service';
import { GoodsReceipt } from '../../../models/GoodsReceipt';

@Component({
    selector: 'app-index-goodsReceipt',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexGoodsReceiptComponent implements OnInit {

    goodsReceipts: GoodsReceipt[] = [];

    constructor(
        private router: Router,
        private service: GoodsReceiptService
) {}

    ngOnInit(): void {
        this.getGoodsReceipts();
}

    getGoodsReceipts(): void {
        this.service.getGoodsReceipts().subscribe((res) => {
        this.goodsReceipts = res;
    });
}

    deleteGoodsReceipt(id: any): void {
        this.service.deleteGoodsReceipt(id)
            .subscribe(() => {
                this.getGoodsReceipts();
            });
    }
}