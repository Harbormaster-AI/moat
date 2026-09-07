
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WorkAuthorizationService } from '../../../services/WorkAuthorization.service';
import { WorkAuthorization } from '../../../models/WorkAuthorization';

@Component({
    selector: 'app-index-workAuthorization',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWorkAuthorizationComponent implements OnInit {

    workAuthorizations: WorkAuthorization[] = [];

    constructor(
        private router: Router,
        private service: WorkAuthorizationService
) {}

    ngOnInit(): void {
        this.getWorkAuthorizations();
}

    getWorkAuthorizations(): void {
        this.service.getWorkAuthorizations().subscribe((res) => {
        this.workAuthorizations = res;
    });
}

    deleteWorkAuthorization(id: any): void {
        this.service.deleteWorkAuthorization(id)
            .subscribe(() => {
                this.getWorkAuthorizations();
            });
    }
}