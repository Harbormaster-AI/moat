
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftOptionComponent } from './index.component';
import { AircraftOptionService } from '../../../services/AircraftOption.service';

describe('IndexAircraftOptionComponent', () => {
  let component: IndexAircraftOptionComponent;
  let fixture: ComponentFixture<IndexAircraftOptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftOptionComponent
      ],
      providers: [
        AircraftOptionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftOptionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});