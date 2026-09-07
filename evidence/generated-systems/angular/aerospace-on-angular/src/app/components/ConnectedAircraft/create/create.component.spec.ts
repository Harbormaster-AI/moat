
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateConnectedAircraftComponent } from './create.component';
import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';
import { Router } from '@angular/router';

describe('CreateConnectedAircraftComponent', () => {
  let component: CreateConnectedAircraftComponent;
  let fixture: ComponentFixture<CreateConnectedAircraftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateConnectedAircraftComponent
      ],
      providers: [
        ConnectedAircraftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateConnectedAircraftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});