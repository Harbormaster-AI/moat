
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BackgroundCheckService } from '../../../services/BackgroundCheck.service';
import { BackgroundCheck } from '../../../models/BackgroundCheck';

@Component({
    selector: 'app-index-backgroundCheck',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBackgroundCheckComponent implements OnInit {

    backgroundChecks: BackgroundCheck[] = [];

    constructor(
        private router: Router,
        private service: BackgroundCheckService
) {}

    ngOnInit(): void {
        this.getBackgroundChecks();
}

    getBackgroundChecks(): void {
        this.service.getBackgroundChecks().subscribe((res) => {
        this.backgroundChecks = res;
    });
}

    deleteBackgroundCheck(id: any): void {
        this.service.deleteBackgroundCheck(id)
            .subscribe(() => {
                this.getBackgroundChecks();
            });
    }
}