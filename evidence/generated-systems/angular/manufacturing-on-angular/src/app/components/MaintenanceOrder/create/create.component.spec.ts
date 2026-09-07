
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMaintenanceOrderComponent } from './create.component';
import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';
import { Router } from '@angular/router';

describe('CreateMaintenanceOrderComponent', () => {
  let component: CreateMaintenanceOrderComponent;
  let fixture: ComponentFixture<CreateMaintenanceOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMaintenanceOrderComponent
      ],
      providers: [
        MaintenanceOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMaintenanceOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});