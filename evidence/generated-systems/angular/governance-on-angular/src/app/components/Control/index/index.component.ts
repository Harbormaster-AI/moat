
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ControlService } from '../../../services/Control.service';
import { Control } from '../../../models/Control';

@Component({
    selector: 'app-index-control',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexControlComponent implements OnInit {

    controls: Control[] = [];

    constructor(
        private router: Router,
        private service: ControlService
) {}

    ngOnInit(): void {
        this.getControls();
}

    getControls(): void {
        this.service.getControls().subscribe((res) => {
        this.controls = res;
    });
}

    deleteControl(id: any): void {
        this.service.deleteControl(id)
            .subscribe(() => {
                this.getControls();
            });
    }
}