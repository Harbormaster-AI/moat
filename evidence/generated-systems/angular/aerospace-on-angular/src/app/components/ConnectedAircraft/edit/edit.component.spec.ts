
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditConnectedAircraftComponent } from './edit.component';
import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';

describe('EditConnectedAircraftComponent', () => {
  let component: EditConnectedAircraftComponent;
  let fixture: ComponentFixture<EditConnectedAircraftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditConnectedAircraftComponent
      ],
      providers: [
        ConnectedAircraftService,
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

    fixture = TestBed.createComponent(EditConnectedAircraftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});