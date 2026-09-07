
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BIQueryService } from '../../../services/BIQuery.service';
import { BIQuery } from '../../../models/BIQuery';

@Component({
    selector: 'app-index-bIQuery',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBIQueryComponent implements OnInit {

    bIQuerys: BIQuery[] = [];

    constructor(
        private router: Router,
        private service: BIQueryService
) {}

    ngOnInit(): void {
        this.getBIQuerys();
}

    getBIQuerys(): void {
        this.service.getBIQuerys().subscribe((res) => {
        this.bIQuerys = res;
    });
}

    deleteBIQuery(id: any): void {
        this.service.deleteBIQuery(id)
            .subscribe(() => {
                this.getBIQuerys();
            });
    }
}