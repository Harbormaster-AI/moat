
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsuredObjectComponent } from './index.component';
import { InsuredObjectService } from '../../../services/InsuredObject.service';

describe('IndexInsuredObjectComponent', () => {
  let component: IndexInsuredObjectComponent;
  let fixture: ComponentFixture<IndexInsuredObjectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsuredObjectComponent
      ],
      providers: [
        InsuredObjectService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsuredObjectComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});