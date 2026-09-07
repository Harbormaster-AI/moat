
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexConnectedAircraftComponent } from './index.component';
import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';

describe('IndexConnectedAircraftComponent', () => {
  let component: IndexConnectedAircraftComponent;
  let fixture: ComponentFixture<IndexConnectedAircraftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexConnectedAircraftComponent
      ],
      providers: [
        ConnectedAircraftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexConnectedAircraftComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});