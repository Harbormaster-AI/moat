
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAircraftComponent } from './create.component';
import { AircraftService } from '../../../services/Aircraft.service';
import { Router } from '@angular/router';

describe('CreateAircraftComponent', () => {
  let component: CreateAircraftComponent;
  let fixture: ComponentFixture<CreateAircraftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAircraftComponent
      ],
      providers: [
        AircraftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAircraftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});