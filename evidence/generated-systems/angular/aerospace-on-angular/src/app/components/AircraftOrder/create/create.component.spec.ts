
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftOrderComponent } from './create.component';
import { AircraftOrderService } from '../../../services/AircraftOrder.service';
import { Router } from '@angular/router';

describe('CreateAircraftOrderComponent', () => {
  let component: CreateAircraftOrderComponent;
  let fixture: ComponentFixture<CreateAircraftOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftOrderComponent
      ],
      providers: [
        AircraftOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});