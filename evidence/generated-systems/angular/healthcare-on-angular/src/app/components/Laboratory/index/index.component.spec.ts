
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLaboratoryComponent } from './index.component';
import { LaboratoryService } from '../../../services/Laboratory.service';

describe('IndexLaboratoryComponent', () => {
  let component: IndexLaboratoryComponent;
  let fixture: ComponentFixture<IndexLaboratoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLaboratoryComponent
      ],
      providers: [
        LaboratoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLaboratoryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});