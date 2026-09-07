
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditMaintenanceOrderComponent } from './edit.component';
import { MaintenanceOrderService } from '../../../services/MaintenanceOrder.service';

describe('EditMaintenanceOrderComponent', () => {
  let component: EditMaintenanceOrderComponent;
  let fixture: ComponentFixture<EditMaintenanceOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditMaintenanceOrderComponent
      ],
      providers: [
        MaintenanceOrderService,
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

    fixture = TestBed.createComponent(EditMaintenanceOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});