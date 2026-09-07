
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMaintenanceWorkOrderComponent } from './index.component';
import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';

describe('IndexMaintenanceWorkOrderComponent', () => {
  let component: IndexMaintenanceWorkOrderComponent;
  let fixture: ComponentFixture<IndexMaintenanceWorkOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMaintenanceWorkOrderComponent
      ],
      providers: [
        MaintenanceWorkOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMaintenanceWorkOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});