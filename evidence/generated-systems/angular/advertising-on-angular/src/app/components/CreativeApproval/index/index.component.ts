
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CreativeApprovalService } from '../../../services/CreativeApproval.service';
import { CreativeApproval } from '../../../models/CreativeApproval';

@Component({
    selector: 'app-index-creativeApproval',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCreativeApprovalComponent implements OnInit {

    creativeApprovals: CreativeApproval[] = [];

    constructor(
        private router: Router,
        private service: CreativeApprovalService
) {}

    ngOnInit(): void {
        this.getCreativeApprovals();
}

    getCreativeApprovals(): void {
        this.service.getCreativeApprovals().subscribe((res) => {
        this.creativeApprovals = res;
    });
}

    deleteCreativeApproval(id: any): void {
        this.service.deleteCreativeApproval(id)
            .subscribe(() => {
                this.getCreativeApprovals();
            });
    }
}