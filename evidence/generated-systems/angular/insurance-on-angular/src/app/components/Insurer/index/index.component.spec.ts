
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInsurerComponent } from './index.component';
import { InsurerService } from '../../../services/Insurer.service';

describe('IndexInsurerComponent', () => {
  let component: IndexInsurerComponent;
  let fixture: ComponentFixture<IndexInsurerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInsurerComponent
      ],
      providers: [
        InsurerService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInsurerComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});