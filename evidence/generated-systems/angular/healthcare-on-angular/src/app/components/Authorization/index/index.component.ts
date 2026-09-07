
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthorizationService } from '../../../services/Authorization.service';
import { Authorization } from '../../../models/Authorization';

@Component({
    selector: 'app-index-authorization',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAuthorizationComponent implements OnInit {

    authorizations: Authorization[] = [];

    constructor(
        private router: Router,
        private service: AuthorizationService
) {}

    ngOnInit(): void {
        this.getAuthorizations();
}

    getAuthorizations(): void {
        this.service.getAuthorizations().subscribe((res) => {
        this.authorizations = res;
    });
}

    deleteAuthorization(id: any): void {
        this.service.deleteAuthorization(id)
            .subscribe(() => {
                this.getAuthorizations();
            });
    }
}