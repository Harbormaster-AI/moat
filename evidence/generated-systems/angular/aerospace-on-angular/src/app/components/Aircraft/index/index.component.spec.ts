
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftComponent } from './index.component';
import { AircraftService } from '../../../services/Aircraft.service';

describe('IndexAircraftComponent', () => {
  let component: IndexAircraftComponent;
  let fixture: ComponentFixture<IndexAircraftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftComponent
      ],
      providers: [
        AircraftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});