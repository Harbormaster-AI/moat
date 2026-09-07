
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditMaintenanceWorkOrderComponent } from './edit.component';
import { MaintenanceWorkOrderService } from '../../../services/MaintenanceWorkOrder.service';

describe('EditMaintenanceWorkOrderComponent', () => {
  let component: EditMaintenanceWorkOrderComponent;
  let fixture: ComponentFixture<EditMaintenanceWorkOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditMaintenanceWorkOrderComponent
      ],
      providers: [
        MaintenanceWorkOrderService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditMaintenanceWorkOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});