
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BOMService } from '../../../services/BOM.service';
import { BOM } from '../../../models/BOM';

@Component({
    selector: 'app-index-bOM',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBOMComponent implements OnInit {

    bOMs: BOM[] = [];

    constructor(
        private router: Router,
        private service: BOMService
) {}

    ngOnInit(): void {
        this.getBOMs();
}

    getBOMs(): void {
        this.service.getBOMs().subscribe((res) => {
        this.bOMs = res;
    });
}

    deleteBOM(id: any): void {
        this.service.deleteBOM(id)
            .subscribe(() => {
                this.getBOMs();
            });
    }
}