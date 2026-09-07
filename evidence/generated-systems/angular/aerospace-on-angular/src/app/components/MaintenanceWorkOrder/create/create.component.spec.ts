
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMaintenanceWorkOrderComponent } from './create.component';
import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';
import { Router } from '@angular/router';

describe('CreateMaintenanceWorkOrderComponent', () => {
  let component: CreateMaintenanceWorkOrderComponent;
  let fixture: ComponentFixture<CreateMaintenanceWorkOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMaintenanceWorkOrderComponent
      ],
      providers: [
        MaintenanceWorkOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMaintenanceWorkOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});