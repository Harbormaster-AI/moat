
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Component_Service } from '../../../services/Component_.service';
import { Component_ } from '../../../models/Component_';

@Component({
    selector: 'app-index-component_',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexComponent_Component implements OnInit {

    component_s: Component_[] = [];

    constructor(
        private router: Router,
        private service: Component_Service
) {}

    ngOnInit(): void {
        this.getComponent_s();
}

    getComponent_s(): void {
        this.service.getComponent_s().subscribe((res) => {
        this.component_s = res;
    });
}

    deleteComponent_(id: any): void {
        this.service.deleteComponent_(id)
            .subscribe(() => {
                this.getComponent_s();
            });
    }
}