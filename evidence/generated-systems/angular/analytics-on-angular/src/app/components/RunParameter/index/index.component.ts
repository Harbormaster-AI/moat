
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RunParameterService } from '../../../services/RunParameter.service';
import { RunParameter } from '../../../models/RunParameter';

@Component({
    selector: 'app-index-runParameter',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRunParameterComponent implements OnInit {

    runParameters: RunParameter[] = [];

    constructor(
        private router: Router,
        private service: RunParameterService
) {}

    ngOnInit(): void {
        this.getRunParameters();
}

    getRunParameters(): void {
        this.service.getRunParameters().subscribe((res) => {
        this.runParameters = res;
    });
}

    deleteRunParameter(id: any): void {
        this.service.deleteRunParameter(id)
            .subscribe(() => {
                this.getRunParameters();
            });
    }
}