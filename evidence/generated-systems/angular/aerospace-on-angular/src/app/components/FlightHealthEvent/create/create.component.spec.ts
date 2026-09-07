
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFlightHealthEventComponent } from './create.component';
import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';
import { Router } from '@angular/router';

describe('CreateFlightHealthEventComponent', () => {
  let component: CreateFlightHealthEventComponent;
  let fixture: ComponentFixture<CreateFlightHealthEventComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFlightHealthEventComponent
      ],
      providers: [
        FlightHealthEventService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFlightHealthEventComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});