
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAircraftProgramComponent } from './index.component';
import { AircraftProgramService } from '../../../services/AircraftProgram.service';

describe('IndexAircraftProgramComponent', () => {
  let component: IndexAircraftProgramComponent;
  let fixture: ComponentFixture<IndexAircraftProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAircraftProgramComponent
      ],
      providers: [
        AircraftProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAircraftProgramComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});