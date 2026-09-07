
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GoodsReceiptLineService } from '../../../services/GoodsReceiptLine.service';
import { GoodsReceiptLine } from '../../../models/GoodsReceiptLine';

@Component({
    selector: 'app-index-goodsReceiptLine',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexGoodsReceiptLineComponent implements OnInit {

    goodsReceiptLines: GoodsReceiptLine[] = [];

    constructor(
        private router: Router,
        private service: GoodsReceiptLineService
) {}

    ngOnInit(): void {
        this.getGoodsReceiptLines();
}

    getGoodsReceiptLines(): void {
        this.service.getGoodsReceiptLines().subscribe((res) => {
        this.goodsReceiptLines = res;
    });
}

    deleteGoodsReceiptLine(id: any): void {
        this.service.deleteGoodsReceiptLine(id)
            .subscribe(() => {
                this.getGoodsReceiptLines();
            });
    }
}